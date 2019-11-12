//  This file is part of zli
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <https://www.gnu.org/licenses/>.

package swap

import "github.com/spf13/cobra"

var api string
var chainId int
var walletAddress string
var gasPrice string
var gasLimit string
var amount string
var priority bool


var SwapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Just for internal swap",
	Long:  "Just for internal swap",
	Run: func(cmd *cobra.Command, args []string) {

	},
}