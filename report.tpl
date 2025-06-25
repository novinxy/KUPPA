<!DOCTYPE html>
<html>

<head>
	<meta charset='utf-8'>
	<meta http-equiv='X-UA-Compatible' content='IE=edge'>
	<title>KUP report</title>
	<meta name='viewport' content='width=device-width, initial-scale=1'>
	<style>
		td {
			padding: 5px;
		}
	</style>
</head>

<body>
	<table>
		<thead>
			<th>Date</th>
			<th>Title</th>
			<th>Message</th>
		</thead>
		<tbody>
            {{range .}}
			<tr>
				<td>{{.Date}}</td>
				<td>{{.Title}}</td>
				<td>{{.Message}}</td>
			</tr>
            {{end}}
		</tbody>
	</table>
</body>

</html>