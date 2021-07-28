import { createPlugin } from '@backstage/core';
import Table from './components/Table';
import User from './components/User';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', User);
    router.registerRoute('/Tables', Table);
  },
});
