import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import {Route, withRouter} from 'react-router-dom';

class Colors extends Component {
    constructor() {
        super();
        this.state = {
            planets: [],
        };
    }

    componentDidMount() {
        let initialPlanets = [];
        fetch("http://localhost:3080/colores")
            .then(response => {
                return response.json();
            }).then(data => {
            initialPlanets = data.results.map((planet) => {
                return planet
            });
            console.log(initialPlanets);
            this.setState({
                planets: initialPlanets,
            });
        });
    }

    render() {
        return (
                    <Planet state={this.state}/>
        );
    }
}

ReactDOM.render(<PlanetSearch />, document.getElementById('react-search'));