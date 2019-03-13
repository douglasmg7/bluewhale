{{define "message"}}
{{ if .HeadMessage }}
<style type="text/css">
  .is-justified-center {
    justify-content: center;
  }
</style>
<div class="message is-medium is-warning is-marginless">
  <div class="message-header is-radiusless is-justified-center">
    <p> {{.HeadMessage}} </p>
  </div>
</div>
{{end}}
{{end}}