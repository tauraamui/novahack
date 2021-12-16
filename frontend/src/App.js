import { Routes, Route } from "react-router-dom";
import { Home } from "./views/Home";
import { Room } from "./views/Room";


export const App = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
          <Route path="/room" element={<Room />} />
        <Route path="/room/:id" element={<h1>Room id: </h1>} />
      </Routes>
    </div>
  );
};
