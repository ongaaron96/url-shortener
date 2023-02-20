import React from 'react';
import { BrowserRouter as Router, Routes, Route, useParams } from 'react-router-dom';
import ReactDOM from 'react-dom/client';
import './index.css';

const CLIENT_DOMAIN_NAME = "http://localhost:3000/"
const SERVER_DOMAIN_NAME = "http://localhost:8081/"
const URL_SHORTENER_PATH = "url-shortener/"
const PROTOCOL_HTTP = "http://"

class App extends React.Component {
  render() {
    return (
      <Router>
        <Routes>
          <Route exact path="/" element={<Home />}></Route>
          <Route exact path="/:shortUrl" element={<Redirect />}></Route>
        </Routes>
      </Router>
    );
  }
}

function Redirect() {
  const { shortUrl } = useParams();
  console.log(shortUrl)

  fetch(SERVER_DOMAIN_NAME + shortUrl, {
    mode: "cors",
    method: "GET",
    headers: { "Content-Type": "application/json" }
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(response.errorMsg);
      }
      return response.json();
    })
    .then(
      (result) => {
        let finalUrl = result.longUrl
        if (!finalUrl.startsWith("http")) {
          finalUrl = PROTOCOL_HTTP + finalUrl
        }
        window.location.href = (finalUrl)
      },
      (error) => {
        console.log(error)
        window.location.href = "/"
      })
    .catch(error => console.log(error))
}

class Home extends React.Component {
  render() {
    return (
      <div className="home">
        <h1>
          Welcome to Aaron's URL Shortener!
        </h1>
        <UrlForm />
      </div>
    )
  }
}

class UrlForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: "",
      isLoaded: false,
      shortUrl: "",
      errorMsg: ""
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(ev) {
    this.setState({
      value: ev.target.value
    });
  }

  handleSubmit(ev) {
    ev.preventDefault();
    this.setState({
      isLoaded: false,
      shortUrl: "",
      errorMsg: ""
    });

    if (this.state.value == "") {
      this.setState({
        errorMsg: "Cannot submit empty URL!"
      });
      return
    }
    fetch(SERVER_DOMAIN_NAME + URL_SHORTENER_PATH + this.state.value, {
      mode: "cors",
      method: "GET",
      headers: { "Content-Type": "application/json" }
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(response.errorMsg);
        }
        return response.json();
      })
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            shortUrl: result.shortUrl,
            errorMsg: ""
          });
        },
        (error) => {
          console.log(error)
          this.setState({
            isLoaded: true,
            errorMsg: "Something went wrong, is the backend server running?"
          });
        })
      .catch(error => console.log(error))
  }

  render() {
    const { isLoaded, shortUrl, error, errorMsg } = this.state
    return (
      <div className="url-form">
        <h2>
          Get your shortened URL here:
        </h2>

        <form onSubmit={this.handleSubmit}>
          <label>
            Your long URL:
            <br /><input type="text" value={this.state.value} onChange={this.handleChange} />
          </label>
          <input type="submit" value="Submit" />
        </form>

        <br />
        {isLoaded && !errorMsg && <div>
          <label>Your short URL: </label>
          <a target="_blank" href={CLIENT_DOMAIN_NAME + shortUrl}>{CLIENT_DOMAIN_NAME + shortUrl}</a>
        </div>}
        {errorMsg && <h3 className="error">{errorMsg}</h3>}
      </div>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<App />);
