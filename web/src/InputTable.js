import React, { useEffect, useState } from 'react';

import ReactDataGrid from 'react-data-grid';

import Tooltip from 'react-bootstrap/Tooltip';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faQuestionCircle } from '@fortawesome/free-solid-svg-icons'

import { Editors, Menu } from "react-data-grid-addons";
import DataContextMenu, { deleteRow, insertRow } from './DataContextMenu';
const { ContextMenuTrigger } = Menu;

export default function InputTable({ title, tooltip, startingRows, defaultRow, cols, onChange }) {
    const [values, setValues] = useState(startingRows);

    const onTableUpdated = ({ fromRow, toRow, updated }) => {
        const r = values.slice();
        for (let i = fromRow; i <= toRow; i++) {
            r[i] = { ...r[i], ...updated };
        }
        setValues(r);

        // Callback
        onChange(r);
    };

    return (
        <div className="col-12">
            <label htmlFor={title.toLowerCase()} className="form-label">{title}</label>
            <OverlayTrigger
                placement="auto"
                overlay={<Tooltip>{tooltip}</Tooltip>}
            >
                <FontAwesomeIcon className="ms-1" icon={faQuestionCircle} />
            </OverlayTrigger>
            <ReactDataGrid
                name={title.toLowerCase()}
                columns={cols}
                rowGetter={i => values[i]}
                rowsCount={values.length}
                onGridRowsUpdated={onTableUpdated}
                enableCellSelect={true}
                minHeight={150}
                contextMenu={
                    <DataContextMenu
                        id={title.toLowerCase() + "ContextMenu"}
                        onRowDelete={(e, { rowIdx }) => setValues(deleteRow(rowIdx, defaultRow))}
                        onRowInsertAbove={(e, { rowIdx }) => setValues(insertRow(rowIdx, defaultRow))}
                        onRowInsertBelow={(e, { rowIdx }) => setValues(insertRow(rowIdx + 1, defaultRow))}
                    />
                }
                RowsContainer={ContextMenuTrigger}
            />
        </div>
    );
}
