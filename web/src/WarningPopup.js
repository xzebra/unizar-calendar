import React from "react";
import { PopupboxManager } from "react-popupbox";
import Popup from './Popup'

function WarningPopup({ warningID }) {
  const { t } = useTranslation();
  return (
    <Popup title={t('warning.title')}>
      <div className="col-12">
        <ReactMarkdown className="line-break">
          {t('warning.' + warningID)}
        </ReactMarkdown>
      </div>
    </Popup>
  );
};

export default function renderWarningPopup(warningID) {
  const content = <WarningPopup
    warningID={warningID}
  />;
  PopupboxManager.open({
    content,
    config: {
      fadeIn: true,
      fadeInSpeed: 400
    }
  })
}
