import ReactDOM from 'react-dom';
import { Route, Switch } from 'react-router';
import { BrowserRouter as Router } from 'react-router-dom';
import { Provider } from 'jotai';

import './styles/base.css';
import './styles/fonts.css';
import './styles/style.css';

import { Home } from './views/Home';
import { Layout, Navbar } from './containers';

const App = () => (
  <Router>
    <Navbar />
    <Switch>
      <Layout>
        {/*Landing page*/}
        <Route exact path="/" component={Home} />
      </Layout>
    </Switch>
  </Router>
);

ReactDOM.render(
  <Provider>
    <App />
  </Provider>,
  document.getElementById('root')
);
