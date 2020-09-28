import React from 'react';

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = { 
        username: '',
        password:'',
        success: '' 
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
            .then(data => this.handleChangeSuccess(data.authenticated));
    }

    handleChangeUsername(event) {
        this.setState({ username: event.target.value });
    }
    handleChangePassword(event) {
        this.setState({ password: event.target.value });
    }

    handleChangeSuccess(sucess_status){
        this.setState({ success: sucess_status });
        if (this.state.success === "true"){
            this.props.handleSetLoggedIn()
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
            .then(data => this.handleChangeSuccess(data.authenticated));
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