import { PC_MIN_WIDTH } from '@site/src/constants/platform';
import useWindow from '@site/src/hooks/useWindow';
import React, { useEffect, useMemo } from 'react';
import Footer from '../components/Footer';
import Header from './header';
import './index.scss';
import Plan from './plan';
import Product from './product';
import Advantage from './advantage';
import Overview from './overview';
import Layout from '@theme/Layout';

export default function Pricing() {
  const { screenWidth } = useWindow();
  const isPc = useMemo(() => screenWidth > PC_MIN_WIDTH, [screenWidth]);

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const params: Record<string, string> = {};

    const bd_vid = urlParams.get('bd_vid');
    if (bd_vid) params.bd_vid = bd_vid;

    const k = urlParams.get('k');
    if (k) params.keywords = k;

    if (Object.keys(params).length > 0) {
      sessionStorage.setItem('sealos_sem', JSON.stringify(params));
    }
  }, []);

  return (
    <div id="sealos-layout-wrap-home-page">
      <Layout>
        <div className="sealos_price_page">
          <img
            draggable="false"
            className="header-img"
            src={require('@site/static/img/license-bg-header.png').default}
            alt="community"
          />
          <div className="px-10 w-full flex justify-center lg:px-0">
            <Header isPc={isPc} />
          </div>
          <Product />
          <Plan />
          <Advantage />
          <Overview />
          <Footer isPc={isPc} />
        </div>
      </Layout>
    </div>
  );
}
