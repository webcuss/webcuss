import React, { useState } from 'react';
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

  switch (route) {
    case Route.Main:
      return (<Main />);
    case Route.About:
      return (<About />);
    case Route.Preference:
      return (<Preference />);
    default:
      return (
        <div>Let's webcuss!</div>
      );
  }
}

export default App;
