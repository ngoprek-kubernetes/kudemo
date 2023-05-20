<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>Untuk Kubernetes Demo yang mudah diikuti.</p>
    <h3>Tautan Penting</h3>
    <ul>
      <li>
        <a href="https://kubernetes.io/id" target="_blank" rel="noopener"
          >Dokumentasi Kubernetes dalam Bahasa Indonesia</a
        >
      </li>
      <br />
      <li>
        <a
          href="https://t.me/kubernetesindonesia"
          target="_blank"
          rel="noopener"
          >Telegram Komunitas Kubernetes Indonesia</a
        >
      </li>
      <br />
      <li>
        <a href="https://community.cncf.io/indonesia/" target="_blank" rel="noopener"
          >Meetup events tentang Cloud Native dan Kubernetes di Indonesia</a
        >
      </li>
    </ul>
    <h3>Penulis</h3>
    <ul>
      <li>
        <a href="https://twitter.com/girikuncoro" target="_blank" rel="noopener"
          >Giri Kuncoro</a
        >
      </li>
      <li>
        <a href="https://twitter.com/ariscahyadi" target="_blank" rel="noopener"
          >Aris Risdianto</a
        >
      </li>
      <li>
        <a href="https://twitter.com/onnowpurbo" target="_blank" rel="noopener"
          >Onno W. Purbo</a
        >
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Hello",
  props: {
    msg: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
