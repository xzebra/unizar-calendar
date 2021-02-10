import React, { Component, useState, useEffect } from 'react';
import CalendarForm from './CalendarForm';
import styled from "styled-components";

const Centered = styled.div`
  display: flex;
  justify-content: center;
  width: 100%;
`

class App extends Component {
  constructor(props) {
    super(props);
  }
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

export default App;
