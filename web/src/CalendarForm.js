import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import ReactDataGrid from 'react-data-grid';
import { Editors, Menu } from "react-data-grid-addons";
import DataContextMenu, { deleteRow, insertRow } from './DataContextMenu';
import styled from "styled-components";

const { DropDownEditor } = Editors;
const { ContextMenuTrigger } = Menu;

const Form = styled.form`
  justify-content: center;
  width: 80%;
`

const defaultColumnProperties = {
  editable: true,
};

const subjectsColumns = [
  { key: 'class_id', name: 'Subject ID' },
  { key: 'class_name', name: 'Subject Name' },
  { key: 'class_desc', name: 'Subject Description' },
].map(c => ({ ...c, ...defaultColumnProperties }));

const subjectsRows = [
  { class_id: 'ia', class_name: 'Inteligencia Artificial', class_desc: 'algo' },
  { class_id: 'ssdd', class_name: 'Sistemas Distribuidos', class_desc: 'otro' },
]

const BoolEditor = <DropDownEditor options={[
  { id: "true", value: "True" },
  { id: "false", value: "False" },
]} />

// weekday;class_id;start_hour;end_hour;is_practical
const schedulesColumns = [
  { key: 'weekday', name: 'Weekday' },
  { key: 'class_id', name: 'Subject ID' },
  { key: 'start_hour', name: 'Start Hour' },
  { key: 'end_hour', name: 'End Hour' },
  { key: 'is_practical', name: 'Is practical', editor: BoolEditor },
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
  const { register, handleSubmit, errors } = useForm();
  const [subjects, setSubjects] = useState(subjectsRows);
  const [schedules, setSchedules] = useState(schedulesRows);
  const [calendarData1, setCalendarData1] = useState("");
  const [calendarData2, setCalendarData2] = useState("");

  useEffect(() => {
    fetch(`/data/semester1.json`)
      .then(response => response.text())
      .then(data => setCalendarData1(data));
    fetch(`/data/semester2.json`)
      .then(response => response.text())
      .then(data => setCalendarData2(data));
  }, []);

  const onSubmit = data => {
    const subjectsData = tableToCSV(subjectsColumns, subjects);
    const schedulesData = tableToCSV(schedulesColumns, schedules);

    // Cast semester to integer so Golang can use it.
    data.semester = parseInt(data.semester);

    console.log(window.calendar(
      data.semester,
      subjectsData,
      schedulesData,
      data.exportType,
      (data.semester === 1 ? calendarData1 : calendarData2),
    ));
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

  return (
    /* "handleSubmit" will validate your inputs before invoking "onSubmit" */
    <Form className="row g-3" onSubmit={handleSubmit(onSubmit)}>
      <div className="col-12">
        <label htmlFor="subjects" className="form-label">Subjects</label>
        <ReactDataGrid
          name="subjects"
          columns={subjectsColumns}
          rowGetter={i => subjects[i]}
          rowsCount={subjects.length}
          onGridRowsUpdated={onSubjectsUpdated}
          enableCellSelect={true}
          minHeight={150}
          contextMenu={
            <DataContextMenu
              id="subjectsContextMenu"
              onRowDelete={(e, { rowIdx }) => setSubjects(deleteRow(rowIdx))}
              onRowInsertAbove={(e, { rowIdx }) => setSubjects(insertRow(rowIdx))}
              onRowInsertBelow={(e, { rowIdx }) => setSubjects(insertRow(rowIdx + 1))}
            />
          }
          RowsContainer={ContextMenuTrigger}
        />
      </div>

      <div className="col-12">
        <label htmlFor="subjects" className="form-label">Schedules</label>
        <ReactDataGrid
          name="schedules"
          columns={schedulesColumns}
          rowGetter={i => schedules[i]}
          rowsCount={schedules.length}
          onGridRowsUpdated={onSchedulesUpdated}
          enableCellSelect={true}
          minHeight={150}
          contextMenu={
            <DataContextMenu
              id="schedulesContextMenu"
              onRowDelete={(e, { rowIdx }) => setSchedules(deleteRow(rowIdx, schedules))}
              onRowInsertAbove={(e, { rowIdx }) => setSchedules(insertRow(rowIdx, schedules))}
              onRowInsertBelow={(e, { rowIdx }) => setSchedules(insertRow(rowIdx + 1, schedules))}
            />
          }
          RowsContainer={ContextMenuTrigger}
        />
      </div>

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
        <input type="submit" className="w-100 btn btn-primary btn-lg" />
      </div>
      <p>{errors.exampleRequired && <span>Some fields are required</span>}</p>
    </Form>
  );
}
