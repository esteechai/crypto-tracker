import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import 'semantic-ui-css/semantic.min.css';
import {StoreContainer} from "./store"

ReactDOM.render(
<StoreContainer.Provider>
<App />
</StoreContainer.Provider>
, document.getElementById('root'));

serviceWorker.unregister();
