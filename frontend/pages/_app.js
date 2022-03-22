//
// Copyright 2021 lemonade
//
import React, { useEffect, useState } from 'react';

import Axios from '../utils/axios';
import Layout from '../components/layout';

function WrappedApp({ Component, pageProps }) {
  const [isLogin, setLogin] = useState(false);
  useEffect(async () => {
    const data = await Axios.get('/session');
    console.log(data);
  }, []);

  return (
    <Layout isLogin={isLogin}>
      <Component {...pageProps} />
    </Layout>
  );
}

export default WrappedApp;
