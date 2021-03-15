import React, { useEffect, Suspense } from 'react';
import ReactDOM from 'react-dom';
import { useTranslation, Trans, I18nextProvider } from 'react-i18next';
import { PopupboxManager, PopupboxContainer } from "react-popupbox";
import styled from "styled-components";

import './i18nextInit';

import CalendarForm from './CalendarForm';
import './index.css';

const Centered = styled.div`
  height: 100%;
  display: flex;
  justify-content: center;
  width: 100%;
`

const Loader = () => {
  <div className="App">
    <div>Loading...</div>
  </div>
}

function App() {
  const { t, i18n } = useTranslation();

  // Run like componentDidMount
  useEffect(() => {
    async function runGolangInstance() {
      // Run golang instance
      const go = new window.Go();
      const source = await fetch(process.env.PUBLIC_URL + "/calendar.wasm");
      const buffer = await source.arrayBuffer();

      let { instance } = await WebAssembly.instantiate(buffer, go.importObject)
      await go.run(instance)
    }

    runGolangInstance();
  }, []);

  return (
    <I18nextProvider i18n={i18n}>
      <Centered>
        <PopupboxContainer />
        <CalendarForm />
      </Centered >
    </I18nextProvider>
  );
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
);
