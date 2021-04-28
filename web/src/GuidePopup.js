import React from "react";

import { useTranslation } from "react-i18next";

import { PopupboxManager } from "react-popupbox";
import ReactMarkdown from 'react-markdown';
import Popup from './Popup'

function GuidePopup({ result, exportType }) {
    const { t } = useTranslation();
    return (
        <Popup title={t('guides.title')}>
            <div className="col-12">
                <ReactMarkdown className="line-break">
                    {t('guides.' + exportType)}
                </ReactMarkdown>
            </div>
        </Popup>
    );
};

export default function renderGuidePopup(res, exportType) {
    const content = <GuidePopup
        result={res}
        exportType={exportType}
    />;
    PopupboxManager.open({
        content,
        config: {
            fadeIn: true,
            fadeInSpeed: 400
        }
    })
}
