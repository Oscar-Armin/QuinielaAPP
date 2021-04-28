import React from 'react'
import Footer from './Footer'
import Navbar from './Navbar'


const Home = ({titulo})=>{
    return(
        <div>
            <div className="Home">
                {titulo}
            </div>
            <Navbar />
            <p> home</p>
            < Footer />
            
        </div>
        
    );
}

export default Home;