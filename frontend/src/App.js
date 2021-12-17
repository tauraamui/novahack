import { Routes, Route } from "react-router-dom";
import { Home } from "./views/Home";
import { Room } from "./views/Room";
import { CaptureImage } from "./views/CaptureImage";
import { Stake } from "./views/Stake";
import { Game } from "./views/Game";

export const App = () => {
  return (
    <div>
      <Routes>
        <Route path="" element={<Home />} />
        <Route path="room/:id" element={<Room />} />
        <Route path="room/:id/game" element={<Game />} />
        <Route path="room/:iod/game/capture" element={<CaptureImage />} />
        <Route path="stake" element={<Stake />} />
      </Routes>
    </div>
  );
};
