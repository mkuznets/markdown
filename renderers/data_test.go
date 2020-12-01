package renderers_test

const testMarkdown = `# Title

**bold** *italic* [url](url)

| one  |   two  |         three |
|------|:------:|--------------:|
| 1    |    4   |             9 |
| один | восемь | двадцать семь |`

const testHTML = `<html>
  <head></head>
  <body>
    <h1>
      Title
    </h1>
    <p>
      <strong>
        bold
      </strong>
      <em>
        italic
      </em>
      <a href="url">
        url
      </a>
    </p>
    <table>
      <thead>
        <tr>
          <th>
            one
          </th>
          <th style="text-align:center">
            two
          </th>
          <th style="text-align:right">
            three
          </th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>
            1
          </td>
          <td style="text-align:center">
            4
          </td>
          <td style="text-align:right">
            9
          </td>
        </tr>
        <tr>
          <td>
            один
          </td>
          <td style="text-align:center">
            восемь
          </td>
          <td style="text-align:right">
            двадцать семь
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>`
