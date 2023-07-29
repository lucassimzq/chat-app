import React, { Component } from "react";

class ChatInput extends Component {
	render() {
		return (
			<div className="px-5">
				<input
					onKeyDown={this.props.send}
					className="border w-full rounded-full shadow px-3 py-2"
				/>
			</div>
		);
	}
}

export default ChatInput;
