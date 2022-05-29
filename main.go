// DONE A. Create Fungsi Login (5 point)
// DONE B. Authorization JWT (6 point)
// DONE C. Reporting nama merchant, omzet per hati dalam bulan november 1 - 30 dengan pagination, if there is no transaction set omzet to 0 (6point)
// DONE D. Buat API untuk reporting nama merchant, nama outlet, omzet perhari pada bulan 1-30 november mulai 1 sampai dengan tanggal 30 dengan pagination. if theres no transaction set omzet to 0 (6 point)
// DONE E. Pada poin C pastikan user tidak bisa melakukan akses pada merchant_id yang bukan miliknya (10 point)
// DONE F.  Pada poin D pastikan user tidak bisa melakukan akses laporan pada outlet_id yang bukan miliknya (5 point)
// TODO G. berikan analisa pdada struktur ERD apakah sudah optimal dengan test case C dan D (9 point)
// TODO H. Buat dokumen teknis Data manipulation language(DML) (3 point)
package main

import "aulia-majoo-test/cmd/api"

func main() {
	api.StartApplication()
}
