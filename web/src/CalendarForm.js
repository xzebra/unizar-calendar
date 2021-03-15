import React, { useEffect, useState } from "react";

import { useForm } from "react-hook-form";
import { isMobile } from "react-device-detect";

import ReactDataGrid from 'react-data-grid';
import { Editors, Menu } from "react-data-grid-addons";
import DataContextMenu, { deleteRow, insertRow } from './DataContextMenu';

import styled from "styled-components";

import "react-popupbox/dist/react-popupbox.css"
import renderErrorPopup from './ErrorPopup'

import fileDownload from 'js-file-download';

import HourEditor from './HourEditor';
import InputTable from './InputTable';

import Tooltip from 'react-bootstrap/Tooltip';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faQuestionCircle } from '@fortawesome/free-solid-svg-icons'

const { DropDownEditor } = Editors;
const { ContextMenuTrigger } = Menu;

const Form = styled.form`
  justify-content: center;
  width: 80%;
`

const defaultColumnProperties = {
  editable: true,
};

const columnTooltipRenderer = (value) => {
  return (
    <span>
      {value.column.name}

      <OverlayTrigger
        placement="auto"
        overlay={<Tooltip>{value.column.tooltip}</Tooltip>}
      >
        <FontAwesomeIcon className="ms-1" icon={faQuestionCircle} />
      </OverlayTrigger>
    </span>
  );
}

const subjectsColumns = [
  {
    key: 'class_id',
    name: 'Subject ID',
    tooltip: 'Subject ID is a shortname of the subject to make it easier to be referenced in Schedules table',
    headerRenderer: columnTooltipRenderer
  },
  { key: 'class_name', name: 'Subject Name' },
  { key: 'class_desc', name: 'Subject Description' },
].map(c => ({ ...c, ...defaultColumnProperties }));

const subjectsRows = [
  { class_id: 'ia', class_name: 'Inteligencia Artificial', class_desc: 'algo' },
  { class_id: 'ssdd', class_name: 'Sistemas Distribuidos', class_desc: 'otro' },
]

const tableTooltip = (isMobile ?
  "Double tap on a cell to edit the content. Hold tap on a row to delete it or insert another one." :
  "Double click on a cell to edit the content. Right click on a row to delete it or insert another one.");

const BoolEditor = <DropDownEditor options={[
  { id: "true", value: "True" },
  { id: "false", value: "False" },
]} />

// weekday;class_id;start_hour;end_hour;is_practical
const schedulesColumns = [
  { key: 'weekday', name: 'Weekday' },
  {
    key: 'class_id',
    name: 'Subject ID',
    tooltip: 'ID specified in Subjects table',
    headerRenderer: columnTooltipRenderer
  },
  { key: 'start_hour', name: 'Start Hour', editor: <HourEditor label="start_hour" /> },
  { key: 'end_hour', name: 'End Hour', editor: <HourEditor label="end_hour" /> },
  {
    key: 'is_practical',
    name: 'Is practical',
    editor: BoolEditor,
    tooltip: 'Some weeks have classes but not practical ones',
    headerRenderer: columnTooltipRenderer
  },
].map(c => ({ ...c, ...defaultColumnProperties }));

const schedulesRows = [
  {
    weekday: 'Lx',
    class_id: 'ia',
    start_hour: '18:00',
    end_hour: '18:50',
    is_practical: "True",
  },
  {
    weekday: 'La',
    class_id: 'ia',
    start_hour: '19:00',
    end_hour: '19:50',
    is_practical: "False",
  },
];

function tableToCSV(columnData, tableData) {
  // I'm sorry if you are reading this, but I was feeling a bit
  // functional today.

  return columnData.map(x => x.key).join(';') +
    '\n' +
    tableData.map(x => Object.entries(x).map(x => x[1]).join(';')).join('\n');
}

export default function CalendarForm() {
  const { register, handleSubmit } = useForm();
  const [subjects, setSubjects] = useState(subjectsRows);
  const [schedules, setSchedules] = useState(schedulesRows);
  const [calendarData1, setCalendarData1] = useState("");
  const [calendarData2, setCalendarData2] = useState("");
  const [result, setResult] = useState("");

  useEffect(() => {
    fetch(process.env.PUBLIC_URL + '/data/semester1.json')
      .then(response => response.text())
      .then(data => setCalendarData1(data));
    fetch(process.env.PUBLIC_URL + '/data/semester2.json')
      .then(response => response.text())
      .then(data => setCalendarData2(data));
  }, []);

  const onSubmit = data => {
    const subjectsData = tableToCSV(subjectsColumns, subjects);
    const schedulesData = tableToCSV(schedulesColumns, schedules);

    // Cast semester to integer so Golang can use it.
    data.semester = parseInt(data.semester);

    const res = window.calendar(
      data.semester,
      subjectsData,
      schedulesData,
      data.exportType,
      (data.semester === 1 ? calendarData1 : calendarData2),
    );

    if (res instanceof Error) {
      renderErrorPopup(res);
      return;
    }

    setResult(res);

    let blob = new Blob([res], {
      type: 'text/plain'
    });
    if (data.exportType === "gcal") {
      fileDownload(blob, "calendar.csv");
    } else if (data.exportType === "org") {
      fileDownload(blob, "calendar.org");
    }

    // renderPopup(res, data.exportType);
  }

  const onSubjectsUpdated = ({ fromRow, toRow, updated }) => {
    const r = subjects.slice();
    for (let i = fromRow; i <= toRow; i++) {
      r[i] = { ...r[i], ...updated };
    }
    setSubjects(r);
  };

  const onSchedulesUpdated = ({ fromRow, toRow, updated }) => {
    const r = schedules.slice();
    for (let i = fromRow; i <= toRow; i++) {
      r[i] = { ...r[i], ...updated };
    }
    setSchedules(r);
  };

  const defaultSubjectsRow = { class_id: '', class_name: '', class_desc: '' };

  const defaultSchedulesRow = {
    weekday: 'Lx',
    class_id: '',
    start_hour: '00:00',
    end_hour: '00:00',
    is_practical: "False",
  };

  return (
    /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
    <Form className="row g-3" onSubmit={handleSubmit(onSubmit)}>
      <InputTable
        title="Subjects"
        tooltip={tableTooltip}
        startingRows={subjectsRows}
        defaultRow={defaultSubjectsRow}
        cols={subjectsColumns}
        onChange={setSubjects}
      />

      <InputTable
        title="Schedules"
        tooltip={tableTooltip}
        startingRows={schedulesRows}
        defaultRow={defaultSchedulesRow}
        cols={schedulesColumns}
        onChange={setSchedules}
      />

      <div className="col-md-6">
        <label htmlFor="semester" className="form-label">Semester</label>
        <select type="number" name="semester" className="form-select" ref={register}>
          <option value={1}>First semester</option>
          <option value={2}>Second semester</option>
        </select>
      </div>

      <div className="col-md-6">
        <label htmlFor="exportType" className="form-label">Export Type</label>
        <select name="exportType" className="form-select" ref={register}>
          <option value="gcal">Google Calendar</option>
          <option value="org">Org Mode</option>
        </select>
      </div>

      <div className="col-md-4">
        <input type="submit"
          className="w-100 btn btn-primary btn-lg" />
      </div>
    </Form >
  );
}
