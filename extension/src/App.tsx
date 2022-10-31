import React, { useEffect, useState } from 'react';
import DebugTabs from './components/debug-tabs/DebugTabs';
import DevToolStorageManager from './components/devtool-storage-manager/DevToolStorageManager';
import PageInfo from './components/page-info/PageInfo';
import { useDebug } from './hooks/useDebug';
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
  const {isDebugging: pIsDebugging} = useDebug();

  const [route, setRoute] = useState<Route>(Route.Main);
  const [activePageIndex, setActivePageIndex] = useState<number>(0);
  const [isDebugging, setIsDebugging] = useState<boolean>(pIsDebugging);

  useEffect(() => {
    setIsDebugging(pIsDebugging);
    if (!pIsDebugging) {
      setActivePageIndex(0);
    }
  }, [pIsDebugging]);

  const debugActiveTabChangeHandler = (tabIndex: number) => {
    setActivePageIndex(tabIndex);
  };

  let rootPage = (<></>);
  switch (activePageIndex) {
    case 0:
      let mainPage = (<></>);
      switch (route) {
        case Route.Main:
          mainPage = (<Main />);
          break;
        case Route.About:
          mainPage = (<About />);
          break;
        case Route.Preference:
          mainPage = (<Preference />);
          break;
        default:
          mainPage = (
            <div>Let's webcuss!</div>
          );
      }
      rootPage = (<>{mainPage}</>);
      break;
    case 1:
      rootPage = (<>
        <PageInfo />
      </>);
      break;
    case 2:
      rootPage = (<>
        <DevToolStorageManager />
      </>);
      break;
  }

  return (
    <>
      {isDebugging && (<DebugTabs onActiveTabChanged={debugActiveTabChangeHandler} />)}
      {rootPage}
    </>
  );
}

export default App;
