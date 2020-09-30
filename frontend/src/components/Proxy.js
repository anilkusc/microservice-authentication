import React from 'react';

class Proxy extends React.Component {
    constructor(props) {
        super(props);
        this.state = {  message: "" , destination: 'https://reqres.in/api/users' , data: '{"name": "morpheus","job": "leader"}'  };
      }
      componentDidMount(){
          console.log(this.state)
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json',"Access-Control-Allow-Origin":"*" },
            body: JSON.stringify({ "destination": this.state.destination,"data": this.state.data})
        };
        fetch('/proxy', requestOptions)
            .then(response => response.json())
            .then(data =>  this.setState({ message: data.name }));
    }
    render() {
        return (
            <div>{this.state.message}</div>
        );
    }
}
export default Proxy;