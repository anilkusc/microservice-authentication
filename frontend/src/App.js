import React from "react";
import Login from "./components/Login"
import Admin from "./components/Admin"
import Index from "./components/Index"
import User from "./components/User"
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
    this.state = {  isLoggedIn: false  , message: "" ,role: "" };
    this.handleSetLoggedIn = this.handleSetLoggedIn.bind(this);
    this.handleUnSetLoggedIn = this.handleUnSetLoggedIn.bind(this);
    this.handleSetRole = this.handleSetRole.bind(this);    
  }

  handleSetRole(UserRole) {
    this.setState({ role: UserRole });
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
    this.setState({ isLoggedIn: false , role: "" });
    
  }
  render(){
    return (
      <Router>
      <div>
        <Links handleUnSetLoggedIn={this.handleUnSetLoggedIn} isLoggedIn={this.state.isLoggedIn} role={this.state.role}/>
        <Switch>
          <PrivateRoute path="/admin" component={ Admin } isLoggedIn={this.state.isLoggedIn}/>
          <PrivateRoute path="/user" component={ User } isLoggedIn={this.state.isLoggedIn}/>
          <PrivateRoute path="/index" component={ Index } isLoggedIn={this.state.isLoggedIn}/>
          <Route path="/login">
          {this.state.isLoggedIn ? <Redirect to="/index" /> : <Login handleSetRole={this.handleSetRole} handleSetLoggedIn={this.handleSetLoggedIn} handleUnSetLoggedIn={this.handleUnSetLoggedIn}/>}
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

class PrivateLink extends React.Component {
   

  render() {
        var link = ""
        if (this.props.role == "user"){
          link =  <li><Link to="/user">User Page</Link></li>
        }else if (this.props.role == "admin"){
          link = <li><Link to="/admin">Admin Page</Link></li>
        }else {
          link = <div></div>
        }        
        return (
        <div>{link}</div>
        );
      
  }
}

class Links extends React.Component {
  isLoggedIn
  
  render() {
    return  this.props.isLoggedIn ? ( <ul>
      <li><Link to="/index">Home Page</Link></li>
        <PrivateLink role={this.props.role}/>
      <li><Link to="/logout" onClick={this.props.handleUnSetLoggedIn}>Logout</Link></li>
    </ul>) : (<div></div>);    
       
      
  }
}