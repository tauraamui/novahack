import { createServer } from "miragejs";

createServer({
  routes() {
    this.namespace = "api";

    this.post("/room", () => {
      const roomId = Math.floor(Math.random() * 1000);
      return roomId;
    });

    this.get("/room", () => {
      return [
        {
          uuid: 1,
          img: "/images/avatar/1.jpg",
          name: "Remy Sharp",
        },
        {
          uuid: 2,
          img: "/images/avatar/2.jpg",
          name: "Travis Howard",
        },
        {
          uuid: 3,
          img: "/images/avatar/3.jpg",
          name: "Cindy Baker",
        },
        {
          uuid: 4,
          img: "/images/avatar/4.jpg",
          name: "Agnes Walker",
        },
        {
          uuid: 5,
          img: "/images/avatar/5.jpg",
          name: "Trevor Henderson",
        },
      ];
    });
  },
});
