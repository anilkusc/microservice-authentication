import React from 'react';
import {
    Link,
    Redirect
  } from "react-router-dom";

class Register extends React.Component {
    constructor(props) {
        super(props);
        this.state = { 
        username: '',
        password:'',
        success: '' ,
        role: '',
        };

        this.handleChangeUsername = this.handleChangeUsername.bind(this);
        this.handleChangePassword = this.handleChangePassword.bind(this);
        this.handleChangeRole = this.handleChangeRole.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChangeSuccess = this.handleChangeSuccess.bind(this);
    }

    handleChangeUsername(event) {
        this.setState({ username: event.target.value });
    }
    handleChangePassword(event) {
        this.setState({ password: event.target.value });
    }
    handleChangeRole(event) {
        this.setState({ role: event.target.value });
    }

    handleChangeSuccess(status){
        this.setState({ success: status });
    }

    handleSubmit(event) {
        event.preventDefault();
        const requestOptions = {
            method: 'POST',
            mode: 'cors',
            headers: { 'Content-Type': 'application/json',"Access-Control-Allow-Origin":"*" },
            body: JSON.stringify({ "username": this.state.username,"password": this.state.password,"role": this.state.role })
        };
        fetch('/register', requestOptions)        
            .then(response => response.json())
            .then(data => this.handleChangeSuccess(data.status));
    }

    render() {
        if (this.state.success === "true") {
            return <Redirect to="/login" />
          }
        return (
            <div>

            <form onSubmit={this.handleSubmit}>
                <br></br>
                <br></br>
                <label>
                    <h3>Register</h3>
                    
                </label><br></br>
                <br></br>
                <label>
                    Username:
                    <input type="text" name="username" onChange={this.handleChangeUsername} />
                </label><br></br>
                <label>
                    Password:
                    <input type="password"  name="password" onChange={this.handleChangePassword} />
                </label>
                <br></br>
                <label>
                    Role:
                    <input type="text"  name="role" onChange={this.handleChangeRole} />
                </label>  
                <br></br>              
                <input type="submit" value="Submit" />
            </form>
            <button ><Link to="/login">Login</Link></button>
            </div>

        );
    }
}
export default Register;