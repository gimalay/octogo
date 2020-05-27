package com.example.todolist.ui.common

import androidx.compose.Composable
import androidx.ui.core.*
import androidx.ui.foundation.shape.corner.RoundedCornerShape
import androidx.ui.graphics.Color
import androidx.ui.layout.Column
import androidx.ui.layout.Padding
import androidx.ui.layout.Row
import androidx.ui.layout.WidthSpacer
import androidx.ui.material.surface.Surface
import androidx.ui.text.TextStyle
import androidx.ui.text.font.FontStyle
import androidx.ui.tooling.preview.Preview

@Composable
fun TextFieldWithHint(
    value: EditorModel,
    label: String = "",
    hint: String = "Please, fill field",
    onChange: (EditorModel) -> Unit
) {
    if (label.isEmpty()) {
        TextFieldOnlyWithHint(
            value = value,
            hint = hint,
            onChange = onChange
        )
    }
    else {
        Row {
            Padding(padding = 8.dp) {
                Text(label)
            }
            WidthSpacer(width = 4.dp)
            TextFieldOnlyWithHint(
                value = value,
                hint = hint,
                onChange = onChange
            )
        }
    }
}

@Composable
fun TextFieldOnlyWithHint(
    value: EditorModel,
    hint: String = "Please, fill field",
    onChange: (EditorModel) -> Unit
) {
    Column {
        Surface(
            shape = RoundedCornerShape(8.dp),
            color = Color.LightGray
        ) {
            Padding(padding = 8.dp) {
                TextField(
                    value = value,
                    onValueChange = onChange
                )
            }
        }
        if (value.text.isEmpty()) {
            Padding(top = 8.dp) {
                Text(
                    text = hint,
                    style = TextStyle(
                        fontStyle = FontStyle.Italic,
                        color = Color.LightGray,
                        fontSize = 12.sp
                    )
                )
            }
        }
    }
}

@Preview
@Composable
fun PreviewTextFieldWithHint() {
    TextFieldWithHint(
        label = "Project Name:",
        value = EditorModel(""),
        hint = "hint",
        onChange = {}
    )
}