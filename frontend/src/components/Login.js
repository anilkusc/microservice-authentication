import React from 'react';

class Login extends React.Component {
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
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChangeSuccess = this.handleChangeSuccess.bind(this);
    }

    componentDidMount(){
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' }
        };
        fetch('/auto', requestOptions)        
            .then(response => response.json())
            .then(data => this.handleChangeSuccess(data));
    }

    handleChangeUsername(event) {
        this.setState({ username: event.target.value });
    }
    handleChangePassword(event) {
        this.setState({ password: event.target.value });
    }

    handleChangeSuccess(auth){
        this.setState({ success: auth.authenticated });
        if (this.state.success === "true"){
            this.props.handleSetLoggedIn()
            this.props.handleSetRole(auth.role)
        }else{
            this.props.handleUnSetLoggedIn()
        }

    }

    handleSubmit(event) {
        event.preventDefault();
        const requestOptions = {
            method: 'POST',
            mode: 'cors',
            headers: { 'Content-Type': 'application/json',"Access-Control-Allow-Origin":"*" },
            body: JSON.stringify({ "username": this.state.username,"password": this.state.password })
        };
        fetch('/login', requestOptions)        
            .then(response => response.json())
            .then(data => this.handleChangeSuccess(data));
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <br></br>
                <br></br>
                <label>
                    Hello {this.state.username}
                    
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
                <input type="submit" value="Submit" />
            </form>
        );
    }
}
export default Login;