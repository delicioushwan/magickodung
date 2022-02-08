//
// Copyright 2021 lemonade
//
import React, { useEffect } from "react";
import Axios from "../utils/axios";

function WrappedApp({ Component, pageProps }) {
  useEffect(async () => {
    await Axios.get("/session");
  }, []);

  return <Component {...pageProps} />;
}

export default WrappedApp;
