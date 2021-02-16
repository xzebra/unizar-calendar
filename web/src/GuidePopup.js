import React from "react";
import fileDownload from 'js-file-download';
import { PopupboxManager } from "react-popupbox";
import Popup from './Popup'

function GuidePopup({ result, guide }) {
    let blob = new Blob([result], {
        type: 'text/plain'
    });

    return (
        <Popup title="How to export">
            <div className="col-12">
                {guide.split("\n").map((i, key) => {
                    return <p key={key}>{i}</p>;
                })}
            </div>
            <div className="col-12">
                <button onClick={() => fileDownload(blob, "test.csv")}>Download</button>
            </div>
        </Popup>
    );
};

export default function renderGuidePopup(res, exportType) {
    const content = <ResultServe
        result={res}
        guide="lorem ipsum"
    />;
    PopupboxManager.open({
        content,
        config: {
            fadeIn: true,
            fadeInSpeed: 400
        }
    })
}
