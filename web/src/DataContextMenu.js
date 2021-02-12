import React from "react";
import { Menu } from "react-data-grid-addons";
const { ContextMenu, MenuItem, SubMenu } = Menu;

export default function DataContextMenu({
    idx,
    id,
    rowIdx,
    onRowDelete,
    onRowInsertAbove,
    onRowInsertBelow
}) {
    return (
        <ContextMenu id={id}>
            <MenuItem data={{ rowIdx, idx }} onClick={onRowDelete}>
                Delete Row
            </MenuItem>
            <SubMenu title="Insert Row">
                <MenuItem data={{ rowIdx, idx }} onClick={onRowInsertAbove}>
                    Above
                </MenuItem>
                <MenuItem data={{ rowIdx, idx }} onClick={onRowInsertBelow}>
                    Below
                </MenuItem>
            </SubMenu>
        </ContextMenu>
    );
}

export const deleteRow = rowIdx => rows => {
    const nextRows = [...rows];
    nextRows.splice(rowIdx, 1);
    return nextRows;
};

export const insertRow = rowIdx => rows => {
    const newRow = {};
    const nextRows = [...rows];
    nextRows.splice(rowIdx, 0, newRow);
    return nextRows;
};
