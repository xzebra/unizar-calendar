import React from "react";
import { PopupboxManager } from "react-popupbox";
import Popup from './Popup'

function ErrorPopup({ err }) {
    return (
        <Popup title="Error">
            {err.split("\n").map((i, key) => {
                return <p key={key}>{i}</p>;
            })}
        </Popup>
    );
};

export default function renderErrorPopup(err) {
    const content = <ErrorPopup
        err={err.message}
    />;
    PopupboxManager.open({
        content,
        config: {
            fadeIn: true,
            fadeInSpeed: 400
        }
    })
}
