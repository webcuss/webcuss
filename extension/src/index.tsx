import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';
import { AuthProvider } from './hooks/useAuth';
import { DebugProvider } from './hooks/useDebug';
import { BrowserExtensionProvider } from './hooks/useBrowserExtension';

const queryClient = new QueryClient();

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserExtensionProvider>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <DebugProvider>
            <App />
          </DebugProvider>
        </AuthProvider>
      </QueryClientProvider>
    </BrowserExtensionProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
