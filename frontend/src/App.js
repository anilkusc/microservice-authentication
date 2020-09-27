import React from "react";
import Login from "./components/Login"
import Admin from "./components/Admin"
import Index from "./components/Index"
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
  Link
} from "react-router-dom";

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {  isLoggedIn: false  , message: ""  };
    this.handleSetLoggedIn = this.handleSetLoggedIn.bind(this);
    this.handleUnSetLoggedIn = this.handleUnSetLoggedIn.bind(this);
  }
  handleSetLoggedIn() {
    this.setState({ isLoggedIn: true });
  }
  handleUnSetLoggedIn() {
    const requestOptions = {
      method: 'POST',
      mode: 'cors',
      headers: { 'Content-Type': 'application/json',"Access-Control-Allow-Origin":"*" },
      body: JSON.stringify({ "username": this.state.username,"password": this.state.password })
  };
  fetch('/login', requestOptions)        
      .then(response => response.json())
      .then(data =>  this.setState({ message: data.authenticated }));
    this.setState({ isLoggedIn: false });
    
  }
  render(){
    return (
      <Router>
      <div>
      <ul>
          <li><Link to="/index">Home Page</Link></li>
          <li><Link to="/admin">Admin Page</Link></li>
          <li><Link to="/login">Login Page</Link></li>
          <li><Link to="/logout" onClick={this.handleUnSetLoggedIn}>Logout</Link></li>
        </ul>
        <Switch>
          <PrivateRoute path="/admin" component={ Admin } isLoggedIn={this.state.isLoggedIn}/>
          <PrivateRoute path="/index" component={ Index } isLoggedIn={this.state.isLoggedIn}/>
          <Route path="/login">
          {this.state.isLoggedIn ? <Redirect to="/index" /> : <Login handleSetLoggedIn={this.handleSetLoggedIn} handleUnSetLoggedIn={this.handleUnSetLoggedIn}/>}
          </Route>
          <Route path="/logout">
          <Redirect to="/login" />
          </Route>
          <Redirect  from="/"  to="/index"  exact  />
        </Switch>
      </div>
    </Router>
   
    );
}

}

class PrivateRoute extends React.Component {
   
  
  render() {
        return  this.props.isLoggedIn ? (<Route  path={this.props.path} component={this.props.component} />) : 
        (<Redirect  to="/login"  />);
      
  }
}
/*
class PublicRoute extends React.Component {
   

  render() {
        return  this.props.isLoggedIn ? (<Route  path={this.props.path} component={this.props.component} />) : 
        (<Redirect  to="/login"  />);
      
  }
}*/