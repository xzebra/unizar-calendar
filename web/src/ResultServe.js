import React from "react";
import fileDownload from 'js-file-download';
import { PopupboxManager } from "react-popupbox";

function ResultServe({ result, guide }) {
    let blob = new Blob([result], {
        type: 'text/plain'
    });

    return (
        <div className="row g-2">
            <div className="col-12">
                {guide.split("\n").map((i, key) => {
                    return <p key={key}>{i}</p>;
                })}
            </div>
            <div className="col-12">
                <button onClick={() => fileDownload(blob, "test.csv")}>Download</button>
            </div>
        </div>
    );
};

export default function renderPopup(res, exportType) {
    const content = <ResultServe
        result={res}
        guide="Idk man"
    />;
    PopupboxManager.open({
        content,
        config: {
            titleBar: {
                enable: true,
                text: 'How to export'
            },
            fadeIn: true,
            fadeInSpeed: 400
        }
    })
}
