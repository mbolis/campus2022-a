import { useState } from "react";

function NewList(props) {
  const [value, setValue] = useState('');

  function submitNewList() {
    if (!value) return;

    props.onNewList(value);
    setValue('');
  }

  return (
    <div>
      <input
        value={value}
        onInput={e => setValue(e.target.value)}
      />
      <button onClick={submitNewList}>+</button>
    </div>
  )
}

export default NewList;
