import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import CalendarForm from './CalendarForm';
import styled from "styled-components";
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
    const source = await fetch("calendar.wasm");

    let { instance } = await WebAssembly.instantiateStreaming(source, go.importObject)
    await go.run(instance)
  }

  render() {
    return (
      <Centered>
        <CalendarForm />
      </Centered>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
);
