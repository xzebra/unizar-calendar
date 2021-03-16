import React, { useEffect, useState } from "react";

import { useForm } from "react-hook-form";
import { isMobile } from "react-device-detect";
import { useTranslation } from "react-i18next";

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

import { Editors } from "react-data-grid-addons";
const { DropDownEditor } = Editors;

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

const BoolEditor = <DropDownEditor options={[
  { id: "true", value: "True" },
  { id: "false", value: "False" },
]} />

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
  const { t } = useTranslation();
  const subjectsColumns = [
    {
      key: 'class_id',
      name: t('tables.subjects.id'),
      tooltip: t('tables.subjects.id_tooltip'),
      headerRenderer: columnTooltipRenderer
    },
    { key: 'class_name', name: t('tables.subjects.name') },
    { key: 'class_desc', name: t('tables.subjects.desc') },
  ].map(c => ({ ...c, ...defaultColumnProperties }));

  const subjectsRows = [
    { class_id: 'ia', class_name: 'Inteligencia Artificial', class_desc: 'algo' },
    { class_id: 'ssdd', class_name: 'Sistemas Distribuidos', class_desc: 'otro' },
  ]

  const tableTooltip = (isMobile ? t('tables.tooltip.mobile') : t('tables.tooltip.web'));

  // weekday;class_id;start_hour;end_hour;is_practical
  const schedulesColumns = [
    { key: 'weekday', name: t('tables.schedules.weekday') },
    {
      key: 'class_id',
      name: t('tables.schedules.subject_id'),
      tooltip: t('tables.schedules.subject_id_tooltip'),
      headerRenderer: columnTooltipRenderer
    },
    { key: 'start_hour', name: t('tables.schedules.start_hour'), editor: <HourEditor label="start_hour" /> },
    { key: 'end_hour', name: t('tables.schedules.end_hour'), editor: <HourEditor label="end_hour" /> },
    {
      key: 'is_practical',
      name: t('tables.schedules.is_practical'),
      editor: BoolEditor,
      tooltip: t('tables.schedules.is_practical_tooltip'),
      headerRenderer: columnTooltipRenderer
    },
  ].map(c => ({ ...c, ...defaultColumnProperties }));

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
        title={t('form.subjects')}
        tooltip={tableTooltip}
        startingRows={subjectsRows}
        defaultRow={defaultSubjectsRow}
        cols={subjectsColumns}
        onChange={setSubjects}
      />

      <InputTable
        title={t('form.schedules')}
        tooltip={tableTooltip}
        startingRows={schedulesRows}
        defaultRow={defaultSchedulesRow}
        cols={schedulesColumns}
        onChange={setSchedules}
      />

      <div className="col-md-6">
        <label htmlFor="semester" className="form-label">{t('form.semester')}</label>
        <select type="number" name="semester" className="form-select" ref={register}>
          <option value={1}>{t('semester.first_semester')}</option>
          <option value={2}>{t('semester.second_semester')}</option>
        </select>
      </div>

      <div className="col-md-6">
        <label htmlFor="exportType" className="form-label">{t('form.export_type')}</label>
        <select name="exportType" className="form-select" ref={register}>
          <option value="gcal">Google Calendar</option>
          <option value="org">Org Mode</option>
        </select>
      </div>

      <div className="col-md-4">
        <input
          type="submit"
          className="w-100 btn btn-primary btn-lg"
          value={t('form.submit')} />
      </div>
    </Form >
  );
}
