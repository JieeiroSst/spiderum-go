import React, { Component } from "react";

import { connect, sendMsg } from "./api";

import Header from "./components/Header/Header";
// import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput";
import Message from "./components/Message";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
    });
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    const messages = this.state.chatHistory.map((msg) => (
      <Message message={msg.data} />
    ));
    return (
      <div className="App">
        <Header />
        {/* <ChatHistory chatHistory={this.state.chatHistory} /> */}
        <div className="ChatHistory">
          {messages}
        </div>
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;
