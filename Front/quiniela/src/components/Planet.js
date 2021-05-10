import React from 'react';

class Planet extends React.Component {
    constructor() {
        super();
    }

    render () {
        let planets = this.props.state.planets;
        let optionItems = planets.map((planet) =>
                <option key={planet.color}>{planet.color}</option>
            );

        return (
         <div>
             <select>
                {optionItems}
             </select>
         </div>
        )
    }
}

export default Planet;