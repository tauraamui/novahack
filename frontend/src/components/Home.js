import { Link } from "react-router-dom";

export const Home = () => {
  return (
    <div>
      <h1>Meta betting</h1>

      <Link to="/room">
        <button>Create new event</button>
      </Link>
    </div>
  );
};
