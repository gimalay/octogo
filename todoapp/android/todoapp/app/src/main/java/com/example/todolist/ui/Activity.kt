package com.example.todolist.ui

import androidx.compose.Composable
import androidx.compose.unaryPlus
import androidx.ui.core.Text
import androidx.ui.core.dp
import androidx.ui.layout.Gravity
import androidx.ui.layout.Row
import androidx.ui.layout.Spacing
import androidx.ui.material.Button
import androidx.ui.material.ContainedButtonStyle
import androidx.ui.material.MaterialTheme
import androidx.ui.material.TextButtonStyle
import androidx.ui.tooling.preview.Preview

@Preview
@Composable
fun Activity(
    text: String,
    onCopy: () -> Unit,
    onRemove: () -> Unit
) {
    Row(
        modifier = Spacing(left = 16.dp, right = 16.dp, top = 5.dp, bottom = 5.dp)
    ) {
        Text(
            text = text,
            modifier = Flexible(1f) wraps Gravity.Center /*wraps Spacing(16.dp)*/,
            style = (+MaterialTheme.typography()).subtitle1
        )

        Row {
            Button(
                "Copy",
                style = ContainedButtonStyle(),
                onClick = onCopy
            )
            Button(
                "Remove",
                style = TextButtonStyle(),
                onClick = onRemove
            )
        }
    }
}