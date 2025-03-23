<?php
$tasksFile = 'tasks.json';
$tasks = file_exists($tasksFile) ? json_decode(file_get_contents($tasksFile), true) : [];

if ($_SERVER['REQUEST_METHOD'] === 'POST' && !empty($_POST['task'])) {
    $newTask = trim($_POST['task']);
    if (!empty($newTask)) {
        $tasks[] = $newTask;
        file_put_contents($tasksFile, json_encode($tasks, JSON_PRETTY_PRINT));
    }
}

if (isset($_GET['delete'])) {
    $index = intval($_GET['delete']);
    if (isset($tasks[$index])) {
        unset($tasks[$index]);
        $tasks = array_values($tasks);
        file_put_contents($tasksFile, json_encode($tasks, JSON_PRETTY_PRINT));
    }
    header("Location: index.php");
    exit;
}
?>
<!DOCTYPE html>
<html>
<head>
    <title>To-Do List</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <h1>To-Do List</h1>
    <form method="POST">
        <input type="text" name="task" placeholder="Enter a task" required>
        <button type="submit">Add Task</button>
    </form>
    <ul>
        <?php foreach ($tasks as $index => $task): ?>
            <li>
                <?= htmlspecialchars($task) ?>
                <a href="?delete=<?= $index ?>">Delete</a>
            </li>
        <?php endforeach; ?>
    </ul>
</body>
</html>