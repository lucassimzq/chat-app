// src/components/Message/Message.jsx
import React, { Component } from "react";

class Message extends Component {
	constructor(props) {
		super(props);
		let temp = JSON.parse(this.props.message);
		this.state = {
			message: temp,
		};
	}

	render() {
		return (
			<div className="bg-white rounded-full shadow inline float-left clear-both px-3 mb-3">
				{this.state.message.body}
			</div>
		);
	}
}

export default Message;
