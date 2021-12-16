import { createServer, Model } from "miragejs";

createServer({
  models: {
    user: Model,
    room: Model,
  },

  seeds(server) {
    server.create("room", { id: "abc123" });
    server.create("user", { uuid: "abc", name: "Joe Doe" });
    server.create("user", { uuid: "cba", name: "Foo Bar" });
    server.create("user", { uuid: "aaa", name: "Ronaldo" });
  },

  routes() {
    this.namespace = "api";

    this.post("/room", () => {
      const roomId = Math.floor(Math.random() * 1000);
      return roomId;
    });

    this.get("/room", (schema) => {
      const participants = schema.user.all();
      return participants;
    });
  },
});
