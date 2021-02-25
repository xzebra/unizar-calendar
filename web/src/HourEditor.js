import React from 'react';
import ReactDOM from "react-dom";
import ReactDataGrid from 'react-data-grid';
import TimeKeeper from 'react-timekeeper';

class HourEditor extends React.Component {
    constructor(props) {
        super(props);
        this.state = { hour: props.value };
        this.label = props.label;
    }

    getValue() {
        return { [this.label]: this.state.hour };
    }

    getInputNode() {
        return ReactDOM.findDOMNode(this).getElementsByTagName("input")[0];
    }

    handleChangeComplete = data => {
        this.setState({ hour: data.formatted24 }, () => this.props.onCommit());
    };

    render() {
        return (
            <TimeKeeper
                time={this.state.hour}
                onChange={this.handleChangeComplete}
            />
        );
    }
}

export default HourEditor;
