import React, { Component } from 'react';
import axios from 'axios';

class Fib extends Component {
  state = {
    seenIndexes: [],
    values: [],
    index: '',
  };

  sleep = ms => new Promise(r => setTimeout(r, ms))

  componentDidMount() {
    this.fetchValues();
    this.fetchIndexes();
  }

  async fetchValues() {
    const values = await axios.get('/api/values/current');
    this.setState({ values: values.data });
  }

  async fetchIndexes() {
    const seenIndexes = await axios.get('/api/values/all');
    this.setState({
      seenIndexes: seenIndexes.data,
    });
  }

  handleSubmit = async (event) => {
    event.preventDefault();

    await axios.post('/api/values/all', {
      index: parseInt(this.state.index, 10),
    });
    await axios.post('/api/values/current', {
      value: parseInt(this.state.index, 10),
    });
    this.setState({ index: '' });

    await this.sleep(500)
    this.componentDidMount()
  };

  renderSeenIndexes() {
    return this.state.seenIndexes.map(({ index }) => index).join(', ');
  }

  renderValues() {
    return this.state.values.map(({key, value}) => (
      <div key={key}>
        For index {key} I calculated {value}
      </div>
    ));
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <label>Enter your index:</label>
          <input
            value={this.state.index}
            onChange={(event) => this.setState({ index: event.target.value })}
          />
          <button>Submit</button>
        </form>

        <h3>Indexes I have seen:</h3>
        {this.renderSeenIndexes()}

        <h3>Calculated Values:</h3>
        {this.renderValues()}
      </div>
    );
  }
}

export default Fib;
