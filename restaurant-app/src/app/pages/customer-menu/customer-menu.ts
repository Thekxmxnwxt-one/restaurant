import { Component, OnInit } from '@angular/core';
import { Menu, MenuService } from '../../services/menu';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-customer-menu',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './customer-menu.html',
  styleUrl: './customer-menu.css'
})
export class CustomerMenu implements OnInit {
  menus: Menu[] = [];
  loading = false;
  errorMessage = '';
  categories: string[] = [];

  constructor(
    private menuService: MenuService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loading = true;
    this.menuService.getMenus().subscribe({
      next: (data: Menu[]) => {
        this.menus = data.filter(menu => menu.available);
        this.categories = [...new Set(data.map(menu => menu.category).filter(c => !!c))];
        this.loading = false;
        console.log(this.menus);
      },
      error: (error: any) => {
        console.error('Menu load error:', error);
        this.errorMessage = 'เกิดข้อผิดพลาดในการโหลดเมนู';
        this.loading = false;
      }
    });
  }

  newMenu: Partial<Menu> = {
    name: '',
    description: '',
    image_url: '',
    price: 0,
    category: '',
    available: true
  };

  addMenu(): void {
    if (!this.newMenu.name || !this.newMenu.description || !this.newMenu.price) {
      alert('กรุณากรอกข้อมูลให้ครบ');
      return;
    }

    this.menuService.createMenu(this.newMenu).subscribe({
      next: (menu) => {
        this.menus.push(menu); // อัปเดตเมนูในหน้าทันที
        alert('เพิ่มเมนูเรียบร้อยแล้ว');
        this.newMenu = { name: '', description: '', image_url: '', price: 0, category: '', available: true };
      },
      error: (err) => {
        console.error('ไม่สามารถเพิ่มเมนูได้:', err);
        alert('เกิดข้อผิดพลาดในการเพิ่มเมนู');
      }
    });
  }

  loadMenus() {
    this.menuService.getMenus().subscribe({
      next: (data) => {
        this.menus = data;
      },
      error: (err) => {
        console.error('โหลดเมนูไม่สำเร็จ:', err);
      }
    });
  }


  editingMenuId: number | null = null;

  startEdit(menu: any) {
    this.editingMenuId = menu.id;
  }

  // ยกเลิก
  cancelEdit() {
    this.editingMenuId = null;
    this.loadMenus(); // โหลดใหม่เพื่อ reset
  }

  // บันทึก
  saveEdit(menu: any) {
    this.menuService.updateMenu(menu.id, menu).subscribe({
      next: () => {
        this.editingMenuId = null;
        this.loadMenus();
      },
      error: () => {
        alert('ไม่สามารถบันทึกการแก้ไขได้');
      }
    });
  }


  deleteMenu(menuId: number): void {
    const confirmed = confirm('คุณต้องการลบเมนูนี้ใช่หรือไม่?');
    if (!confirmed) return;

    this.menuService.deleteMenu(menuId).subscribe({
      next: () => {
        this.menus = this.menus.filter(menu => menu.id !== menuId);
        alert('ลบเมนูเรียบร้อยแล้ว');
      },
      error: (err) => {
        console.error('ลบเมนูล้มเหลว:', err);
        alert('ไม่สามารถลบเมนูได้');
      }
    });
  }
}
