package com.example.todolist.ui.common

import androidx.compose.Composable
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.core.*
import androidx.ui.foundation.Dialog
import androidx.ui.foundation.shape.corner.RoundedCornerShape
import androidx.ui.graphics.Color
import androidx.ui.layout.*
import androidx.ui.material.*
import androidx.ui.material.surface.Surface
import androidx.ui.tooling.preview.Preview
import com.example.todolist.ui.lightThemeColors
import com.example.todolist.ui.themeTypography

@Composable
fun DialogButton(
    visible: Boolean = false,
    disabled: Boolean = false,
    text: String,
    dialogTitle: String = text,
    onApply: () -> Unit = {},
    children: @Composable() () -> Unit
) {
    val visibleModel = +state { visible }

    Button(
        text,
        style = ContainedButtonStyle(),
        onClick = {
            visibleModel.value = true
        }
    )
    if (visibleModel.value) {
        DialogForm(
            title = dialogTitle,
            disabled = disabled,
            onClose = {
                visibleModel.value = false
            },
            onApply = {
                onApply()
                visibleModel.value = false
            }
        ) {
            children()
        }
    }
}

@Composable
fun DialogForm(
    title: String,
    disabled: Boolean,
    onApply: () -> Unit,
    onClose: () -> Unit,
    children: @Composable() () -> Unit
) {
    Dialog(onCloseRequest = onClose) {
        MaterialTheme(
            colors = lightThemeColors,
            typography = themeTypography
        ) {
            Surface(shape = RoundedCornerShape(4.dp)) {
                Container {
                    Column {
                        DialogTitle(title)
                        children()
                        HeightSpacer(height = 28.dp)
                        DialogActions(
                            disabled = disabled,
                            onApply = onApply,
                            onClose = onClose
                        )
                    }
                }
            }
        }
    }
}

@Composable
fun DialogTitle(titleText: String) {
    Container(
        alignment = Alignment.CenterLeft,
        padding = EdgeInsets(
            left = 24.dp, top = 24.dp, right = 24.dp, bottom = 0.dp
        )
    ) {
        val textStyle = (+MaterialTheme.typography()).h6
        CurrentTextStyleProvider(textStyle) {
            Text(titleText)
        }
    }
}

@Composable
fun DialogActions(
    disabled: Boolean,
    onApply: () -> Unit,
    onClose: () -> Unit
) {
    Container(
        ExpandedWidth,
        padding = EdgeInsets(all = 8.dp),
        alignment = Alignment.CenterRight
    ) {
        Row {
            if (disabled) {
                Button(
                    "Apply",
                    style = ContainedButtonStyle().copy(
                        color = Color.LightGray
                    )
                )
            } else {
                Button(
                    "Apply",
                    onClick = onApply
                )
            }
            WidthSpacer(8.dp)
            Button(
                "Dismiss",
                style = ContainedButtonStyle().copy(
                    color = Color.Transparent
                ),
                onClick = onClose
            )
        }
    }
}

@Preview
@Composable
fun PreviewDialogActions() {
    Column {
        DialogActions(disabled = true, onApply = {}, onClose = {})
        DialogActions(disabled = false, onApply = {}, onClose = {})
    }
}