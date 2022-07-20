import { useEffect, useState } from 'react';
import './App.css';

import List from './components/List';
import NewList from './components/NewList';

function App() {
  const [lists, setLists] = useState([]);

  useEffect(() => {
    fetch('/lists')
      .then(resp => resp.json())
      .then(result => setLists(result.lists))
      .catch(console.error);
  }, []);

  return (
    <div className="App">
      {lists.map(l => (
        <List
          key={l.id}
          title={l.title}
        />
      ))}
      <NewList onNewList={title => {
        fetch('/lists', {
          method: 'POST',
          body: JSON.stringify({title}),
        })
          .then(resp => resp.json())
          .then(result => setLists([...lists, result]))
          .catch(console.error);
      }} />
    </div>
  );
}

export default App;
