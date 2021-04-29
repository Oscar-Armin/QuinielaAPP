
import React from "react";
import Footer from "./Footer";
import Navbar from "./Navbar";

const LoadImage = ({titulo})=>{
    return(
<>  
        <Navbar />
        <form
      enctype="multipart/form-data"
      action="http://localhost:3080/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
        <Footer />
</>
    );
}

export default LoadImage;