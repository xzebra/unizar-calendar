import React, { useEffect, useState } from 'react';

import ReactDataGrid from 'react-data-grid';

import Tooltip from 'react-bootstrap/Tooltip';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faQuestionCircle } from '@fortawesome/free-solid-svg-icons'

import { Editors, Menu } from "react-data-grid-addons";
import DataContextMenu, { deleteRow, insertRow } from './DataContextMenu';
const { ContextMenuTrigger } = Menu;

export default function InputTable({ title, tooltip, startingRows, defaultRow, columns, onChange }) {
    const [values, setValues] = useState(startingRows);
    const [canvas, setCanvas] = useState(null);
    const [context, setContext] = useState(null);

    const onTableUpdated = ({ fromRow, toRow, updated }) => {
        const r = values.slice();
        for (let i = fromRow; i <= toRow; i++) {
            r[i] = { ...r[i], ...updated };
        }
        setValues(r);

        let newCols = formatColumns(columns, r);
        console.log(newCols);
        setCols(newCols);

        // Callback
        onChange(r);
    };

    const formatColumns = (cols, values) => {
        // const gridWidth = parseInt(document.querySelector("#root").clientWidth, 10); //selector for grid
        let combinedColumnWidth = 0;

        for (let i = 0; i < cols.length; i++) {
            cols[i].width = getTextWidth(cols, values, i);
            combinedColumnWidth += cols[i].width;
        }

        // if (combinedColumnWidth < gridWidth) {
        //     data.columns = distributeRemainingSpace(
        //         combinedColumnWidth,
        //         data.columns,
        //         gridWidth
        //     );
        // }
        return cols;
    }

    const getTextWidth = (cols, values, i) => {
        const rowValues = [];
        const reducer = (a, b) => (a.length > b.length ? a : b);
        const cellPadding = 16;
        const arrowWidth = 18;
        let longestCellData,
            longestCellDataWidth,
            longestColName,
            longestColNameWidth,
            longestString;

        for (let row of values) {
            rowValues.push(row[cols[i].key]);
        }

        longestCellData = rowValues.reduce(reducer);
        longestColName = cols[i].name;
        longestCellDataWidth = Math.ceil(
            getCanvas().measureText(longestCellData).width
        );
        longestColNameWidth =
            Math.ceil(getCanvas("bold ").measureText(longestColName).width) +
            arrowWidth;

        longestString = Math.max(longestCellDataWidth, longestColNameWidth);

        return longestString + cellPadding;
    };

    const getCanvas = (fontWeight = "") => {
        let ctx = context;

        if (!canvas) {
            let c = document.createElement("canvas");
            setCanvas(c);
            ctx = c.getContext("2d")
            setContext(ctx);
        }

        ctx.font = `${fontWeight}24px`;

        return ctx;
    };

    const [cols, setCols] = useState(formatColumns(columns, values));

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
