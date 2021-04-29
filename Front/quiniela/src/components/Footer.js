import React from 'react'

function Footer(){
    return(
        <footer className="bg-light text-center text-lg-start">
        <div className="main-footer">
            <div className="container">
                <div className="footer-bottom">
                    <center>
                <p className="text-xs-center">
                    &copy;{new Date().getFullYear()} - All Rights Reserved 201709140

                </p>
                </center>
                </div>
            </div>
            
        </div>
        </footer>
    )
}

export default Footer;