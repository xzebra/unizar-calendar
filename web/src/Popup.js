import React from "react";
import { PopupboxManager } from "react-popupbox";

export default function Popup({ title, children }) {
    return (
        <div className="row g-2">
            <div className="d-flex justify-content-between col-12">
                <div>
                    <h2 className="popup-title">{title}</h2>
                </div>
                <div className="d-flex">
                    <button
                        type="button"
                        className="btn-close my-auto"
                        aria-label="Close"
                        onClick={PopupboxManager.close}
                    >
                    </button>
                </div>
            </div>
            <div className="col-12">
                {children}
            </div>
        </div >
    );
};
