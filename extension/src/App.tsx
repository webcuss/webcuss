import React, { useState } from 'react';
import styled from 'styled-components';
import DevToolStorageManager from './components/devtool-storage-manager/DevToolStorageManager';
import About from './pages/about/About';
import Main from './pages/main/Main';
import Preference from './pages/preference/Preference';

enum Route {
  Main,
  SignIn,
  About,
  Preference,
}

const App = () => {
  const [route, setRoute] = useState<Route>(Route.Main);

  let page = (<></>);

  switch (route) {
    case Route.Main:
      page = (<Main />);
      break;
    case Route.About:
      page = (<About />);
      break;
    case Route.Preference:
      page = (<Preference />);
      break;
    default:
      page = (
        <div>Let's webcuss!</div>
      );
  }

  return (
    <>
    <div>main | page-info | storage</div>
      {page}

      <DevTools>
        <DevToolStorageManager />
      </DevTools>
    </>
  );
}

export default App;

const DevTools = styled.div`
  height: 100px;
  border-top: 1px dotted black;
  overflow-y: auto;
`;
