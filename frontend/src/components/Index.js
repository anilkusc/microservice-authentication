import React from 'react';

class Index extends React.Component {
    constructor(props) {
        super(props);
        this.state = {  message: ""  };
      }
      componentDidMount(){
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json',"Access-Control-Allow-Origin":"*" }
        };
        fetch('/index', requestOptions)        
            .then(response => response.json())
            .then(data => this.setState({ message: data.role }));
    }
    render() {
        return (
            <div>{this.state.message}</div>
        );
    }
}
export default Index;