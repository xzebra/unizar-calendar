import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import CalendarForm from './CalendarForm';
import styled from "styled-components";
import { PopupboxManager, PopupboxContainer } from "react-popupbox";
import './index.css';

const Centered = styled.div`
  height: 100%;
  display: flex;
  justify-content: center;
  width: 100%;
`

class App extends Component {
  async componentDidMount() {
    // Run golang instance
    const go = new window.Go();
    const source = await fetch(process.env.PUBLIC_URL + "/calendar.wasm");
    const buffer = await source.arrayBuffer();

    let { instance } = await WebAssembly.instantiate(buffer, go.importObject)
    await go.run(instance)
  }

  render() {
    return (
      <Centered>
        <PopupboxContainer />
        <CalendarForm />
      </Centered>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
);
