import './App.css';
import TaskList from './components/TaskList';

function App() {
  return (
    <>
      <section className="Header">
        Hello, i am a header
      </section>
      <div className="App">
        <TaskList>

        </TaskList>
      </div>
      <section className="Footer">
        Hello, i am a footer
      </section>
    </>
  );
}

export default App;
